## Step 1 - Setting up the Go Application

Today I initialized the Go module and created the initial project structure. Instead of placing everything in a single file, I chose a structure that separates the application entry point (`cmd`) from the internal business logic (`internal`). Although this project is small, this layout is commonly used in production Go applications because it scales well as the project grows.

### Learned

- How Go modules work
- Why production Go projects use `cmd/` and `internal/`
- How to create a basic HTTP server using Go's standard library

Health endpoints allow orchestration platforms and monitoring tools to verify that an application is running correctly without relying on the main user-facing page.

Version endpoints expose the running application version, making deployments easier to verify and troubleshoot.

I implemented graceful shutdown so the application can terminate cleanly when it receives SIGINT or SIGTERM. Instead of abruptly stopping, the server gets up to five seconds to finish in-flight requests before exiting. This behavior is essential for deployments on platforms like Docker and Kubernetes, where applications are routinely stopped and restarted.

Goroutines – Run the server without blocking the rest of the program.
Channels – Receive operating system signals.
Contexts – Give the server up to 5 seconds to finish ongoing work.
Graceful shutdown – Stop accepting new requests while allowing current requests to complete.

Infrastructure as Code allows cloud infrastructure to be defined using code instead of manual portal operations. This makes infrastructure reproducible, version-controlled, and easier to review. Terraform maintains a state file that tracks deployed resources, enabling it to determine what changes are required without recreating everything.
