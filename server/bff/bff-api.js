const typeDefs = require('./schema/schema');
const resolvers = require('./resolver/resolver');

const { RESTDataSource } = require('apollo-datasource-rest');
const { ApolloServer } = require('apollo-server');

// 各APIに対するリクエスト処理を記述する
class UserAPI extends RESTDataSource {
  constructor() {
    super();
    this.baseURL = process.env.USER;
  }

  async setUser(phoneNum) {
    return this.post(
        `set`,
        {phone: phoneNum},
      );
  }
}

// graphQLの初期化
const server = new ApolloServer({
  typeDefs,
  resolvers,
  dataSources: () => {
    return {
        // MicroServiceに対するリクエスト処理のインスタンスを追加
        UserAPI: new UserAPI(),
    };
  },
});

server.listen().then(({ url }) => {
  console.log(`Server ready at ${url}`);
});