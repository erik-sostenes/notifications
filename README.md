# notifications-api

_REST API service in charge of managing real-time notifications each time an action takes place_

## Note: the architectural pattern Hexagonal Architecture (ports and adapters was used)

## Hexagonal Architecture

The hexagonal architecture is based on three principles and techniques:
1.Explicitly separate User-Side, Business Logic, and Server-Side
2.Dependencies are going from User-Side and Server-Side to the Business Logic
3.We isolate the boundaries by using Ports and Adapters

```
├───cmd                         👉🏼 (execute commands)
│   └───bootstrap               👉🏼 (bootstrap package that builds the program with its full set of components)
├───internal
│   ├───core                    👉🏼 (core business)
│   │   ├───module              👉🏼 (represents a boundary)
│   │   │   ├───business        👉🏼 (business logic layer)
│   │   │   │   ├───domain      👉🏼 (data transfer objects, business objects, errors, entities, value object)
│   │   │   │   ├───ports       👉🏼 (business contracts)
│   │   │   │   └───services    👉🏼 (logic)
│   │   │   └───infrastructure  👉🏼 (layer infrastructure)
│   │   │       ├───driven      👉🏼 (output adapters)
│   │   │       └───drives      👉 (input adapters)
```

### User-Side

- This is the side through which the user or external programs will interact with the application.

### Business Logic

- This is the part that we want to isolate from both left and right sides. It contains all the code that concerns and implements business logic.

### Server-Side

- This is where we’ll find what your application needs, what it drives to work. It contains essential infrastructure details such as the code that interacts with your database, makes calls to the file system, or code that handles HTTP calls to other applications on which you depend for example.

```

    |------------|                   |----------------|             |-------------|
    | User side  | =====[port]=====> | Business logic | <==[port]== | server side |
    |------------|                   |----------------|             |-------------|

```
