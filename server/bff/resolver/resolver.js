module.exports = {
    Query: {
        // リクエストに応じたクエリを作成
        setUser: async (_source, { phone },{ dataSources }) => {
            return dataSources.UserAPI.setUser(phone);
        },
        setRate: async (_source, { id, rate },{ dataSources }) => {
            return dataSources.RateAPI.setRate(id, rate);
        },
        setPresent: async (_source, { id, asin },{ dataSources }) => {
            return dataSources.PresentAPI.setPresent(id, asin);
        },
        setMessage: async (_source, { id, message },{ dataSources }) => {
            return dataSources.SendMessageAPI.setMessage(id, message);
        },
    },
};