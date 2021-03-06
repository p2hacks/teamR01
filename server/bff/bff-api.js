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

class RateAPI extends RESTDataSource {
  constructor() {
    super();
    this.baseURL = process.env.RATE;
  }

  async setRate(userId, userRate) {
    return this.post(
        `set`,
        { id: userId, rate: userRate },
      );
  }
}

class PresentAPI extends RESTDataSource {
  constructor() {
    super();
    this.baseURL = process.env.PRESENT;
  }

  async setPresent(userId, asin) {
    return this.post(
        `set`,
        { id: userId, ASIN: asin },
      );
  }
}

class GetPairAPI extends RESTDataSource {
  constructor() {
    super();
    this.baseURL = process.env.PAIR;
  }

  async getPair(userId) {
    return this.post(
        `set`,
        { id: userId },
      );
  }
}

class SetCardInfoAPI extends RESTDataSource {
  constructor() {
    super();
    this.baseURL = process.env.CARD_INFO;
  }

  async setCardInfo(id, num, year, month, code, name) {
    return this.post(
        `set`,
        { id: id, num: num, year: year, month: month, code: code, name: name },
      );
  }
}

class SendMessageAPI extends RESTDataSource {
  constructor() {
    super();
    this.baseURL = process.env.SEND_MESSAGE;
  }

  async setMessage(userId, msg) {
    return this.post(
        `set`,
        { id: userId, message: msg },
      );
  }
}

class GetMessageAPI extends RESTDataSource {
  constructor() {
    super();
    this.baseURL = process.env.GET_MESSAGE;
  }

  async getMessage(id) {
    return this.post(
        `set`,
        { id: id },
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
        RateAPI: new RateAPI(),
        PresentAPI: new PresentAPI(),
        SendMessageAPI: new SendMessageAPI(),
        GetMessageAPI: new GetMessageAPI(),
        GetPairAPI: new GetPairAPI(),
        SetCardInfoAPI: new SetCardInfoAPI(),
    };
  },
});

server.listen().then(({ url }) => {
  console.log(`Server ready at ${url}`);
});