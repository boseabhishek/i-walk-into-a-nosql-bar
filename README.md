# i-walk-into-a-nosql-bar


## I visit the AWS DynamoDB corner

I want to write some cool stuff using DynamoDB. Hence, I want a nice but easy scenario.

Hence, I pick up [this]( https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/GSI.html#GSI.scenario) scenario to start with..


### Acceptance criteria

Feature: User tracks gamers and their scores for a mobile gaming application

 - Scenario: User requests scores by gamer's id
 
        - Given there are multiple gamers with game title they have played
            - And the respective scores

        - When I query by gamer's id

        - Then I could see the games they have played
            - And the respective scores.     
