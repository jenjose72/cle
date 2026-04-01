#include <mpi.h>
#include <stdio.h>
#include <unistd.h>

/*
    MPI HELLO WORLD (C)
    This program demonstrates distributed computing. 
    When run across multiple nodes, each process will print its unique rank 
    and the hostname of the machine it is running on.
*/

int main(int argc, char** argv) {
    // Initialize the MPI environment
    MPI_Init(&argc, &argv);

    // Get the number of processes
    int world_size;
    MPI_Comm_size(MPI_COMM_WORLD, &world_size);

    // Get the rank (ID) of the current process
    int world_rank;
    MPI_Comm_rank(MPI_COMM_WORLD, &world_rank);

    // Get the name of the processor (hostname)
    char processor_name[MPI_MAX_PROCESSOR_NAME];
    int name_len;
    MPI_Get_processor_name(processor_name, &name_len);

    // Print off a hello world message
    printf("Hello from rank %d out of %d processes, running on host: %s\n",
           world_rank, world_size, processor_name);

    // Finalize the MPI environment.
    MPI_Finalize();
    
    return 0;
}
