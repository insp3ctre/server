import React from 'react';

import Page from '../pageFrame/Page';
import MyPlaylistList from '../components/lists/MyPlaylistList';

export default class MyPlaylistsContainer extends React.PureComponent {
    render() {
        return <Page component={MyPlaylistList} />
    }
}