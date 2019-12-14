module.exports = {
    Query: {
        // リクエストに応じたクエリを作成
        setUser: async (_source, { phone },{ dataSources }) => {
            return dataSources.UserAPI.setUser(phone);
        },
        setRate: async (_source, { id, rate },{ dataSources }) => {
            return dataSources.RateAPI.setRate(id, rate);
        },
    },
};