# Distributed Computing with VMs and MPI

This guide covers setting up a distributed computing environment using Ubuntu VMs (local and AWS) and executing an MPI application across them.

---

## Task 1: Distributed MPI on Local Ubuntu VMs

### 1. Setup Two Ubuntu VMs (VirtualBox)
1.  **Create VMs**: In VirtualBox, create two Ubuntu 22.04 LTS VMs.
2.  **Networking**: 
    - Settings -> Network -> Adapter 1.
    - Attached to: **Bridged Adapter** (This allows VMs to get their own IP on your local network).
3.  **Hostnames**: Set unique hostnames (e.g., `master-node` and `worker-node`).

### 2. Configure Networking
On **Both** VMs, find their IP addresses (`ip addr`) and edit `/etc/hosts`:
```bash
sudo nano /etc/hosts
# Add entries for both nodes
192.168.x.x master-node
192.168.x.y worker-node
```

### 3. Setup Passwordless SSH (Master to Worker)
On `master-node`:
1.  **Generate Key**: `ssh-keygen -t rsa` (Press Enter through all prompts).
2.  **Copy Key**: `ssh-copy-id worker-node`.
3.  **Verify**: `ssh worker-node` (Should log in without a password).

### 4. Install OpenMPI
On **Both** VMs:
```bash
sudo apt update
sudo apt install -y openmpi-bin libopenmpi-dev build-essential
```

### 5. Compile and Execute
1.  **Compile** on `master-node`:
    ```bash
    mpicc hello_mpi.c -o hello_mpi
    ```
2.  **Distribute**: Ensure the `hello_mpi` executable is in the same path on both machines (e.g., `/home/ubuntu/`).
3.  **Run**:
    ```bash
    mpirun -np 4 --host master-node,worker-node ./hello_mpi
    ```

---

## Task 2: Create a Ubuntu VM on Amazon AWS

1.  **Login**: Go to the [AWS Management Console](https://aws.amazon.com/console/).
2.  **EC2**: Navigate to the EC2 Dashboard and click **Launch Instance**.
3.  **Configuration**:
    - **Name**: `Ubuntu-Server`
    - **AMI**: Ubuntu Server 22.04 LTS (HVM), SSD Volume Type.
    - **Instance Type**: `t2.micro` (Free Tier Eligible).
    - **Key Pair**: Create or select an existing `.pem` key pair.
4.  **Network Settings**:
    - Enable **Allow SSH traffic from** (Select "My IP" for security).
5.  **Launch**: Click **Launch Instance**.
6.  **Connect**:
    ```bash
    chmod 400 your-key.pem
    ssh -i "your-key.pem" ubuntu@<Public-IPv4-DNS>
    ```

---

## 🚀 Expected MPI Output
When running the `hello_mpi` across two nodes, you should see:
```text
Hello from rank 0 out of 4 processes, running on host: master-node
Hello from rank 1 out of 4 processes, running on host: master-node
Hello from rank 2 out of 4 processes, running on host: worker-node
Hello from rank 3 out of 4 processes, running on host: worker-node
```
*(Distribution depends on the number of slots available on each node).*
