const { gql } = require('apollo-server');

// クエリを追加する
module.exports = gql`
  type User {
    STATUS:  Boolean
    ID: String
  }
  type Rate {
    STATUS: Boolean
  }
  type Present {
    STATUS: Boolean!
  }
  type SendMessage {
    STATUS: Boolean!
  }
  type Query {
    setUser(phone: String!): User
    setRate(id: String!, rate: Int!): Rate
    setPresent(id: String, asin: String!): Present
    setMessage(id: String!, message: String): SendMessage
  }
`;