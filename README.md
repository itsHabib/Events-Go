# Events - Go
Sample project used to help practice and learn creating microservices with Go

## Overview
The application has an events microservice used for querying and adding events
to a database. The events microservice itself has a persistence layer, configuration layer, and
serves request using RESTful handlers.The booking service has the same structure and is used for 
booking events. To allow different types of database to be used with  the handlers each potential database handler must implement a defined interface located in lib/persistence/persistence.go The dblayer package 
acts as a middle man and uses a factory pattern to return a database handler given a configuration. 

## Components
1. BookingService
2. EventsService
3. AMQP Implementation
4. Apache Kafka Implementation
5. MongoDB Implementation

## Security
The main file is set up to serve both HTTP & HTTPS requests for the eventservice. Each ListenAndServe call for both HTTP & HTTPS are ran in a go routine and both use channels to send back any errors to the main go routine.

## Task List
- [ ] complete events microservice
- [x] complete booking microservice
- [x] add MongoDB db handler implementation
- [ ] add DynamoDB db handler implementation
- [x] add message queue interface to allow for multiple implementations
- [x] add amqp implementation
- [x] add kafka implementation
- [ ] add React.js front end