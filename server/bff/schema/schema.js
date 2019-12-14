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
  type GetMessage {
    STATUS: Boolean!
    MESSAGE: String
  }
  type GetPair {
    STATUS: Boolean!
    ASIN: String
  }
  type SetCard {
    STATUS: Boolean!
  }
  type Query {
    setUser(phone: String!): User
    setRate(id: String!, rate: Int!): Rate
    setPresent(id: String, asin: String!): Present
    setMessage(id: String!, message: String): SendMessage
    getMessage(id: String!): GetMessage
    getPair(id: String!): GetPair
    setCardInfo(id: String!, , num: String!, year: String!, month: String!, code: String!, name: String!): SetCard!
  }
`;