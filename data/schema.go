package data

import (
  "github.com/graphql-go/graphql"
  "github.com/graphql-go/relay"
)

var operationType *graphql.Object
var operationList *graphql.Object
var queryType *graphql.Object
var Schema graphql.Schema

func init() {

  operationType = graphql.NewObject(graphql.ObjectConfig{
    Name: "Operation",
    Fields: graphql.Fields{
      // Define `id` field as a Relay GlobalID field.
      // It helps with translating your GraphQL object's id into a global id
      // For eg:
      //  For a `Operation` type, with an id of `1`, it's global id will be `UG9zdDox`
      //  which is a base64 encoded version of `Operation:1` string
      // We will explore more in the next part of this series.
      "id": relay.GlobalIDField("Operation", nil),
      "title": &graphql.Field{
        Type: graphql.String,
      },
    },
  })

  operationList = graphql.NewObject(graphql.ObjectConfig{
    Name: "OperationList",
    Fields: graphql.Fields{
      // Define `id` field as a Relay GlobalID field.
      // It helps with translating your GraphQL object's id into a global id
      // For eg:
      //  For a `Operation` type, with an id of `1`, it's global id will be `UG9zdDox`
      //  which is a base64 encoded version of `Operation:1` string
      // We will explore more in the next part of this series.
      "id": relay.GlobalIDField("OperationList", nil),
      "operations": &graphql.Field{
        Type: graphql.NewList(operationType),
      },
    },
  })

  queryType = graphql.NewObject(graphql.ObjectConfig{
    Name: "Query",
    Fields: graphql.Fields{
      "getOperation": &graphql.Field{
        Type: operationType,
        Resolve: func(p graphql.ResolveParams) (interface{}, error) {
          return GetOperation(), nil
        },
      },
      "allOperations": &graphql.Field{
        Type: operationList,
        Resolve: func(p graphql.ResolveParams) (interface{}, error) {
          return GetAllOperations(), nil
        },
      },
    },
  })

  Schema, _ = graphql.NewSchema(graphql.SchemaConfig{
    Query: queryType,
  })

}
