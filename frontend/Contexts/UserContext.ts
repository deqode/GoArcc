import { createContext } from "react";

export const UserContext = createContext({
    user: {
        accessToken: "",
        idToken: "",
        userId: ""
    },
    setUser: (value) => { },
    removeUser: () => { }
});
