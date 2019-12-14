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
        getMessage: async (_source, { id },{ dataSources }) => {
            return dataSources.GetMessageAPI.getMessage(id);
        },
        getPair: async (_source, { id },{ dataSources }) => {
            return dataSources.GetPairAPI.getPair(id);
        },
        setCardInfo: async (_source, { id, num, year, month, code, name },{ dataSources }) => {
            return dataSources.SetCardInfoAPI.setCardInfo(id, num, year, month, code, name);
        },
    },
};