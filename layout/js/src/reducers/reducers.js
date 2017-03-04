import Immutable from 'immutable';

import * as defaultStates from './defaultStates';
import * as actionNames from '../actions/actionNames';

export function mySongs(state = defaultStates.MY_SONGS, action) {
    switch (action.type) {
        case actionNames.SET_MY_SONGS:
            return action.payload;
        default:
            return state;
    }
}

export function myArtists(state = defaultStates.MY_ARTISTS, action) {
    switch (action.type) {
        case actionNames.SET_MY_ARTISTS:
            return action.payload;
        default:
            return state;
    }
}

export function myAlbums(state = defaultStates.MY_ALBUMS, action) {
    switch (action.type) {
        case actionNames.SET_MY_ALBUMS:
            return action.payload;
        default:
            return state;
    }
}

export function play(state = defaultStates.PLAY, action) {
    switch (action.type) {
        case actionNames.PLAY_SONG:
            return Immutable.List.of(action.payload);
        default:
            return state;
    }
}