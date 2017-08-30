import React, { Component } from 'react'
import { Collapse, message } from 'antd';
import {injectIntl, intlShape, FormattedMessage} from 'react-intl'

import Layout from '../../../layouts/Dashboard'
import {get} from '../../../ajax'

const Panel = Collapse.Panel;

class WidgetF extends Component {
  state = {
    items: [],
  }
  componentDidMount () {
    get('/survey/reports').then(
      function (rst){
        this.setState(rst)
      }.bind(this)).catch(message.error)
  }
  render() {
    const zone = (k, o) => (<Panel header={<FormattedMessage id={`site.admin.status.${k}`}/>} key={k}>
      <table width="100%">
        <tbody className="ant-table-tbody">
          {Object.entries(o).map((v, i)=><tr key={i}><td>{v[0]}</td><td>{v[1]}</td></tr>)}
        </tbody>
      </table>
    </Panel>)
    return (
      <Layout admin breadcrumbs={[{href: '/survey/reports', label: <FormattedMessage id='survey.reports.index.title'/>}]}>
        <Collapse defaultActiveKey={[]}>
          
        </Collapse>
      </Layout>
    );
  }
}


WidgetF.propTypes = {
  intl: intlShape.isRequired,
}

export default injectIntl(WidgetF)
