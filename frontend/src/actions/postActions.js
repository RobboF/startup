import { FETCH_SITES, NEW_SITE } from './types'

export const fetchPosts = () => dispatch => {
    console.log(process.env.REACT_APP_API_URL)
    fetch(`${process.env.REACT_APP_API_URL}/items`)
        .then(res => res.json())
        .then(items => 
            dispatch({
                type: FETCH_SITES,
                payload: items
            })
        )
}

export const createPost = postData => dispatch => {
    fetch('http://api/item', {
        method: 'POST',
        headers: {'Content-Type': 'application/json'},
        body: JSON.stringify(postData)
    })
    .then(res => res.json())
    .then(item => dispatch({
        type: NEW_SITE,
        payload: item
    }))
}