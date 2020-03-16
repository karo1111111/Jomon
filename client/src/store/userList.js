import axios from "axios";
import Vue from "vue";
export const userList = {
  state: [
    {
      trap_id: "user1",
      is_admin: true
    },
    {
      trap_id: "user2",
      is_admin: false
    },
    {
      trap_id: "user3",
      is_admin: false
    },
    {
      trap_id: "user4",
      is_admin: false
    }
  ],
  mutations: {
    setUserList(state, newState) {
      state.splice(0);
      newState.forEach((element, index) => {
        Vue.set(state, index, element);
      });
    }
  },
  actions: {
    async getUserList({ commit }) {
      try {
        const response = await axios.get("/api/users");
        commit("setUserList", response.data);
      } catch (err) {
        console.log(err);
      }
    }
  }
};
