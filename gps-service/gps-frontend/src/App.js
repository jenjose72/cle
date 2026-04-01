import { MapContainer, TileLayer, Marker, Popup } from "react-leaflet";
import { useEffect, useState } from "react";
import axios from "axios";
import "leaflet/dist/leaflet.css";

function App() {

  const [location, setLocation] = useState(null);
  const [userLocation, setUserLocation] = useState(null);
  const [sendStatus, setSendStatus] = useState(null);

  // Backend base URL (can be overridden with REACT_APP_BACKEND_URL)
  const BACKEND_URL = process.env.REACT_APP_BACKEND_URL || "http://localhost:5000";

  useEffect(() => {
    // fetch current server-side stored location (classroom or last submitted)
    axios.get(`${BACKEND_URL}/location`)
      .then(res => {
        setLocation(res.data);
      })
      .catch(err => console.warn("Could not fetch location from backend:", err.message));

    // ask for browser geolocation and send it to the backend
    if (navigator.geolocation) {
      navigator.geolocation.getCurrentPosition(async (pos) => {
        const coords = pos.coords;
        const payload = {
          latitude: coords.latitude,
          longitude: coords.longitude,
          timestamp: new Date().toISOString()
        };

        setUserLocation(payload);
        setSendStatus("sending");
        try {
          await axios.post(`${BACKEND_URL}/location`, payload);
          setSendStatus("sent");
          // refresh displayed location to reflect latest
          const res = await axios.get(`${BACKEND_URL}/location`);
          setLocation(res.data);
        } catch (err) {
          console.error("Failed to send location:", err.message || err);
          setSendStatus("failed");
        }
      }, (err) => {
        console.warn("Geolocation error:", err.message);
        setSendStatus("denied");
      }, { enableHighAccuracy: true, timeout: 10000 });
    } else {
      console.warn("Browser does not support geolocation");
      setSendStatus("unsupported");
    }
  }, []); // run once on mount

  if (!location) return <h2>Loading classroom location...</h2>;

  return (
    <div style={{height:"100vh"}}>
      <MapContainer
        center={[location.latitude, location.longitude]}
        zoom={18}
        style={{height:"100%"}}
      >

        <TileLayer
          url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png"
        />

        <Marker position={[location.latitude, location.longitude]}>
          <Popup>
            <b>2023 Batch 2 - AA103</b><br/>
            Classroom Location
          </Popup>
        </Marker>

      </MapContainer>
      {/* UI for user location status */}
      <div style={{position:"absolute", right:12, top:12, background:"rgba(255,255,255,0.9)", padding:10, borderRadius:6}}>
        <div><b>Location status</b></div>
        <div>Backend: {BACKEND_URL}</div>
        <div>User: {userLocation ? `${userLocation.latitude.toFixed(6)}, ${userLocation.longitude.toFixed(6)}` : 'not captured'}</div>
        <div>Send: {sendStatus || 'idle'}</div>
      </div>
    </div>
  );
}

export default App;