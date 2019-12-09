import { FETCH_SITES, NEW_SITE } from '../actions/types'

const initialState = {
    items: [],
    item: {}
}
export default (state = initialState, action) => {
    switch(action.type) {
        case FETCH_SITES:
            return {
                ...state,
                items: action.payload
            }
        case NEW_SITE:
                return {
                    ...state,
                    item: action.payload
                }
        default:
            return state;            
    }
}