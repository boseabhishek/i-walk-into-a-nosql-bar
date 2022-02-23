# i-walk-into-a-nosql-bar

## My plan

I have picked up [this](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/GSI.html#GSI.scenario) scenario to start with. I want to work on this scenario and see how I can achive this using different dbs e.g. Dynamodb, Couchbase, MongoDB.

Below are a list of features I want to work on:

## Features

> Feature: User tracks gamers and their scores for a mobile gaming application

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

As per AWS DynamoDB docs,

> NoSQL database systems like Amazon DynamoDB use alternative models for data management, such as key-value pairs or document storage. When you switch from a relational database management system to a NoSQL database system like DynamoDB, it's important to understand the key differences and specific design approaches.

I want to implement the above scenarios using DynamoDB.

### some dynamodb basics:

- **partition key**

  The primary key that uniquely identifies each item in an Amazon DynamoDB table can be simple (a partition key only) or composite (a partition key combined with a sort key).
  Also known as its hash attribute. The term "hash attribute" derives from the DynamoDB usage of an internal hash function to evenly distribute data items across partitions, based on their partition key values.

- **sort key** :point_up:

  The sort key of an item is also known as its range attribute. The term "range attribute" derives from the way DynamoDB stores items with the same partition key physically close together, in sorted order by the sort key value.

- indexes
  - primary
  - secondary - LSI and GSI
