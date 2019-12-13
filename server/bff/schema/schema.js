const { gql } = require('apollo-server');

// クエリを追加する
module.exports = gql`
  type User {
    STATUS:  Boolean
    ID: String
  }
  type Query {
    setUser(phone: String!): User
  }
`;