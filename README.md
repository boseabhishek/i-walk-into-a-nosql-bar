# i-walk-into-a-nosql-bar


## Dynamodb

Loosely based on https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/GSI.html#GSI.scenario

### Acceptance criteria

Feature: User tracks gamers and their scores for a mobile gaming application

 - Scenario: User requests scores by gamer's id
 
        - Given there are multiple gamers with game title they have played
            - And the respective scores

        - When I query by gamer's id

        - Then I could see the games they have played
            - And the respective scores.     
