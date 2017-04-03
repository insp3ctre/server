import React from 'react';

import Page from '../pageFrame/Page';
import Overview from '../components/overview/Overview';

export default class StartpageContainer extends React.PureComponent {
    render() {
        return <Page component={Overview} />
    }
}