// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.

import CredimintCredimint from './credimint.credimint'
import CredimintLqs from './credimint.lqs'


export default { 
  CredimintCredimint: load(CredimintCredimint, 'credimint.credimint'),
  CredimintLqs: load(CredimintLqs, 'credimint.lqs'),
  
}


function load(mod, fullns) {
    return function init(store) {        
        if (store.hasModule([fullns])) {
            throw new Error('Duplicate module name detected: '+ fullns)
        }else{
            store.registerModule([fullns], mod)
            store.subscribe((mutation) => {
                if (mutation.type == 'common/env/INITIALIZE_WS_COMPLETE') {
                    store.dispatch(fullns+ '/init', null, {
                        root: true
                    })
                }
            })
        }
    }
}