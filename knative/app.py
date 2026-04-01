import math
import re
import time
from collections import Counter

from flask import Flask, jsonify, request

app = Flask(__name__)

POSITIVE_TERMS = {
    "resolved",
    "healthy",
    "stable",
    "recovered",
    "success",
    "fixed",
    "mitigated",
    "normal",
}

NEGATIVE_TERMS = {
    "down",
    "outage",
    "failed",
    "critical",
    "breach",
    "timeout",
    "error",
    "degraded",
    "panic",
    "urgent",
    "latency",
}

STOPWORDS = {
    "the",
    "and",
    "for",
    "with",
    "from",
    "that",
    "this",
    "into",
    "your",
    "about",
    "have",
    "been",
    "after",
    "when",
    "were",
    "across",
    "over",
    "under",
    "will",
    "would",
    "could",
    "there",
    "their",
    "while",
    "where",
    "which",
    "our",
    "are",
    "was",
    "is",
    "in",
    "on",
    "to",
    "of",
    "at",
    "it",
    "an",
    "a",
}


def clamp(value, low=0.0, high=1.0):
    return max(low, min(high, value))


def extract_entities(text):
    ips = re.findall(r"\b(?:\d{1,3}\.){3}\d{1,3}\b", text)
    emails = re.findall(r"\b[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Za-z]{2,}\b", text)
    tickets = re.findall(r"\b(?:INC|SR|TASK)-\d{3,10}\b", text)

    services = []
    for token in re.findall(r"\b[a-zA-Z][a-zA-Z0-9_-]{2,}\b", text):
        lowered = token.lower()
        if lowered in {"kafka", "redis", "postgres", "mysql", "nginx", "gateway", "auth"}:
            services.append(lowered)

    return {
        "ips": sorted(set(ips)),
        "emails": sorted(set(emails)),
        "tickets": sorted(set(tickets)),
        "services": sorted(set(services)),
    }


def extract_key_phrases(text, limit=5):
    words = re.findall(r"\b[a-zA-Z]{3,}\b", text.lower())
    filtered = [w for w in words if w not in STOPWORDS]
    common = Counter(filtered).most_common(limit)
    return [word for word, _ in common]


def sentiment_score(text):
    words = re.findall(r"\b[a-zA-Z]{3,}\b", text.lower())
    if not words:
        return 0.5

    pos_hits = sum(1 for w in words if w in POSITIVE_TERMS)
    neg_hits = sum(1 for w in words if w in NEGATIVE_TERMS)

    raw = (pos_hits - neg_hits) / max(1, len(words) / 8)
    return round(clamp((raw + 1) / 2), 3)


def urgency_score(text, severity_hint=None):
    lower = text.lower()
    keyword_boost = 0.0
    for token in ["critical", "outage", "breach", "payment", "downtime", "production", "p0", "p1"]:
        if token in lower:
            keyword_boost += 0.1

    length_signal = clamp(math.log(max(len(text), 2), 10) / 3)
    severity_signal = {"low": 0.2, "medium": 0.5, "high": 0.75, "critical": 0.95}.get(
        (severity_hint or "").lower(),
        0.4,
    )

    score = clamp((keyword_boost * 0.5) + (length_signal * 0.2) + (severity_signal * 0.3))
    return round(score, 3)


def recommend_owner(entities, urgency):
    services = set(entities.get("services", []))
    if "gateway" in services or "nginx" in services:
        return "platform-networking"
    if "auth" in services:
        return "identity-security"
    if urgency >= 0.8:
        return "incident-command"
    if "postgres" in services or "mysql" in services:
        return "data-platform"
    return "sre-core"


@app.get("/")
def home():
    return jsonify(
        {
            "service": "AIOps Incident Triage API",
            "status": "running",
            "version": "2.0.0",
            "endpoints": ["/health", "/analyze", "/batch-analyze"],
        }
    )


@app.get("/health")
def health():
    return jsonify({"ok": True, "timestamp": int(time.time())})


@app.post("/analyze")
def analyze_incident():
    started = time.perf_counter()
    payload = request.get_json(silent=True) or {}

    text = str(payload.get("text", "")).strip()
    incident_id = str(payload.get("incident_id", "auto-generated"))
    severity_hint = payload.get("severity", None)

    if not text:
        return jsonify({"error": "Field 'text' is required."}), 400

    entities = extract_entities(text)
    sentiment = sentiment_score(text)
    urgency = urgency_score(text, severity_hint)
    owner = recommend_owner(entities, urgency)
    phrases = extract_key_phrases(text)

    elapsed_ms = round((time.perf_counter() - started) * 1000, 2)

    return jsonify(
        {
            "incident_id": incident_id,
            "analysis": {
                "urgency": urgency,
                "sentiment": sentiment,
                "priority": "P1" if urgency >= 0.8 else "P2" if urgency >= 0.6 else "P3",
                "recommended_owner": owner,
                "key_phrases": phrases,
                "entities": entities,
            },
            "runtime": {
                "inference_time_ms": elapsed_ms,
                "model": "hybrid-heuristic-v2",
            },
        }
    )


@app.post("/batch-analyze")
def batch_analyze():
    payload = request.get_json(silent=True) or {}
    incidents = payload.get("incidents", [])

    if not isinstance(incidents, list):
        return jsonify({"error": "Field 'incidents' must be a list."}), 400

    results = []
    for item in incidents:
        text = str(item.get("text", "")).strip() if isinstance(item, dict) else ""
        if not text:
            continue

        entities = extract_entities(text)
        urgency = urgency_score(text, item.get("severity"))
        results.append(
            {
                "incident_id": item.get("incident_id", "auto-generated"),
                "urgency": urgency,
                "owner": recommend_owner(entities, urgency),
                "key_phrases": extract_key_phrases(text, limit=3),
            }
        )

    return jsonify({"count": len(results), "results": results})


if __name__ == "__main__":
    app.run(host="0.0.0.0", port=8080)