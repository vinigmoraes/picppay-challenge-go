# Description

This project is my solution for the challenge of [Picpay Backend Challenge]() storage at open source repository [Backend Challenges]()

## The Development Solution

It's possible to registry a user (Common or Seller) when the user is created a wallet is automatically created for him with 100 balance.

When a user transfer money for another user its published a event on MessageBroker asking for a consumer to execute this money transfer, this consumer receive the event to transfer money for the receiver wallet.

In this consumer has a integration with a authorizer

- if the authorize respond successful the money of payer is transfer to the receiver.
- If not the consumer publish a message to DLQ Queue to give money back to the payer.

## Running Application


## Application Design

The organization of the project was struct based on [Standard Go Project Design](https://github.com/golang-standards/project-layout), inside the basic structure I choose to use Hexagonal Architecture.

## Architecture


## Utility Links

The content above its in portuguese from Brazil 🇧🇷

- [How pointer works in Golang]()
- [How to organize your Golang Project](https://www.youtube.com/watch?v=OFud4iPuAH8) 


## Built With

- [GinGonic](https://gin-gonic.com/) - Web Framework
- [GoORM](https://gorm.io/index.html) - Database ORM framework
- [RabbitMQ]() - Message Broker
- [Docker](https://www.docker.com/) - Containerization Platform