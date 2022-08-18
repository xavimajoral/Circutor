
export async function GetBookmarks() {

    const response = await fetch(`/user/bookmarks`, {
        method: 'GET',
        headers: { 'Content-Type': 'application/json', 'Authorization': `Bearer ${localStorage.getItem('test-token')}` }
    })

    return await response.json();

}

export async function SaveBookmarks(_id) {

    const response = await fetch(`/user/bookmarks`, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json', 'Authorization': `Bearer ${localStorage.getItem('test-token')}` },
        body: JSON.stringify({ buildingId: _id })
    })

    return await response.json();

}

export async function DeleteBookmarks(_id) {

    const response = await fetch(`/user/bookmarks/${_id}`, {
        method: 'DELETE',
        headers: { 'Content-Type': 'application/json', 'Authorization': `Bearer ${localStorage.getItem('test-token')}` },
    })

    return await response.json();

}