Secure and Scalable Golang Microservices with Docker Integration.
This repo is a Golang microservices framework designed for secure, scalable development. It utilizes a hexagonal architecture and dependency injection for loose coupling and 
enhanced testability. 

Communication between services is a combination of:
gRPC: Provides high-performance, remote procedure calls (RPC) for efficient service interaction.
RabbitMQ: Enables asynchronous, event-driven communication, allowing services to react to events.

Modular Services for Secure User Management:
user: This service handles core user functionalities like registration, user information retrieval, and potentially other user-related operations.
auth: Focused on secure authentication, this service handles logins, generates access tokens, and validates credentials.
authorize: This service manages authorization logic, allowing for role assignment and permission checks for specific resources within the system.

Flexible Data Persistence:
offers flexibility in data persistence by supporting both:
PostgreSQL: A powerful, open-source object-relational database (ORDBMS) for complex data models.
MySQL: A widely used, open-source relational database management system (RDBMS) for simpler data structures.

Streamlined Development Workflow with Docker:
Image Building: gir leverages Docker for building containerized images for each microservice, ensuring consistent environments and simplifying deployments.
Orchestration: Docker Compose or other orchestration tools can be used to manage and scale deployments of gir's microservices on a container platform like Docker Swarm or Kubernetes.

This repository provides a robust foundation for building secure, scalable, and maintainable microservices in Golang, with the added benefit of a streamlined development 
workflow using Docker for image building and orchestration.
