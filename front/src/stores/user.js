import { writable } from 'svelte/store'


const user = writable({
  token: null,
  firstName: null,
  lastName: null,
})

export default user
