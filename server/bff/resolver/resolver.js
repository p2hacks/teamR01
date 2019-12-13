module.exports = {
    Query: {
        // リクエストに応じたクエリを作成
        setUser: async (_source, { phone },{ dataSources }) => {
            return dataSources.UserAPI.setUser(phone);
        },
    },
};