# i-walk-into-a-nosql-bar

## My plan

> This might be a blog post someday 🎨

I have picked up [this](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/GSI.html#GSI.scenario) scenario to start with. I want to work on this scenario and see how I can achive this using different dbs e.g. Dynamodb, Couchbase, MongoDB. The main objective is to use `Go` to write e2e scnearios(some might call it integration as we are using local docker images as infra) but my opinion is e2e as I am testing business cenarios.

The main purpose is to benchmark using [Go's benchmarking](https://dave.cheney.net/2013/06/30/how-to-write-benchmarks-in-go) to study the performance of each db techs, in terms of:

- use of secondary index

Below are a list of features I want to work on:

## Features

### Feature: User tracks gamers and their scores for a mobile gaming application

> Scenario: User requests scores by gamer's id and game title

        - Given there are multiple gamers with game title they have played
            - And the respective scores

        - When I query by gamer's id and game title

        - Then I could fetch the top score for that gamer

> Scenario: User requests topscore of each game

        - Given there are multiple gamers with game title they have played
            - And the respective scores

        - When I query by game title

        - Then I could fetch the top score for that game title

> Scenario: User requests gemer's id

        - Given there are multiple gamers with game title they have played
            - And the respective scores

        - When I query by game title

        - Then I could fetch the top score for that game title        

## :beer: first stop: AWS DynamoDB lounge

> For a DynamoDb 101, see [here](docs/dynamodb.md)

I want to implement the above scenarios using DynamoDB.

What do I find which can help me writing future code when it comes to working with dynamodb?

_**data retrieval**_

- Single-item requests (PutItem, GetItem, UpdateItem, and DeleteItem) that act on a single, specific item and require the full primary key;
- Query, which can read a range of items and must include the partition key;
- Scan, which can read a range of items but searches across your entire table.

Other than the wildly-inefficient (for most use cases!) Scan operation, you must include the partition key in any request to DynamoDB.


## TODOs:

- CI

## Resources:

- [SQL to NoSQL](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/SQLtoNoSQL.html)

