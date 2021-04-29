import { createContext } from "react";

export const UserContext = createContext({
    user: {
        accessToken: "",
        idToken: "",
        userId: "",
        provider:""
    },
    setUser: (value) => { },
    removeUser: () => { }
});
