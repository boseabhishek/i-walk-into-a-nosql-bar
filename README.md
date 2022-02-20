# i-walk-into-a-nosql-bar

## My plan

 I have picked up [this]( https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/GSI.html#GSI.scenario) scenario to start with. I want to work on this scenario and see how I can achive this using different dbs e.g. Dynamodb, Couchbase, MongoDB.

## 🎉 I visit the AWS DynamoDB corner

As per AWS DynamoDB docs,

>NoSQL database systems like Amazon DynamoDB use alternative models for data management, such as key-value pairs or document storage. When you switch from a relational database management system to a NoSQL database system like DynamoDB, it's important to understand the key differences and specific design approaches.

I want to write some cool stuff using DynamoDB. 
Also, I want to explore the below concepts:

- **partition key**

    The primary key that uniquely identifies each item in an Amazon DynamoDB table can be simple (a partition key only) or composite (a partition key combined with a sort key).

- **sort key** :point_up:

- indexes
    - primary
    - secondary - LSI and GSI


### Acceptance criteria

Feature: User tracks gamers and their scores for a mobile gaming application

 - Scenario: User requests scores by gamer's id
 
        - Given there are multiple gamers with game title they have played
            - And the respective scores

        - When I query by gamer's id

        - Then I could see the games they have played
            - And the respective scores.     
