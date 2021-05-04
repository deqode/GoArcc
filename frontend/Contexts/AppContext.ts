import Centrifuge from "centrifuge";
import { createContext } from "react";
import { CENTRIFUGO } from '../utils/constants';

export const AppContext = createContext({
    app: {
        centrifuge: new Centrifuge(CENTRIFUGO),
        subscribed: false,
    },
    setApp: (value) => {},
    removeApp: () => {}
});
