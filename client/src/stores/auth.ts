import {writable} from 'svelte/store';

export const authToken = writable(localStorage.getItem('authToken') || '');

authToken.subscribe(value => {
    localStorage.setItem('authToken', value);
});