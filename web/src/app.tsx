// Polyfill URL because Chrome and Firefox are not spec-compliant
// Hostnames of URIs with custom schemes (e.g. git) are not parsed out
import './util/polyfill'

import ErrorIcon from '@sourcegraph/icons/lib/Error'
import ServerIcon from '@sourcegraph/icons/lib/Server'
import * as React from 'react'
import { render } from 'react-dom'
import { Route, RouteComponentProps, Switch } from 'react-router'
import { BrowserRouter } from 'react-router-dom'
import { Subscription } from 'rxjs/Subscription'
import { fetchCurrentUser } from './auth'
import { HeroPage } from './components/HeroPage'
import { updateUserSessionStores } from './marketing/util'
import { Navbar } from './nav/Navbar'
import { routes } from './routes'
import { parseSearchURLQuery } from './search'
import { InitializePage } from './settings/InitializePage'
import { colorTheme, getColorTheme } from './settings/theme'

interface LayoutProps extends RouteComponentProps<any> {
    isLightTheme: boolean
}

const Layout: React.SFC<LayoutProps> = props => {
    const isSearchHomepage = props.location.pathname === '/search' && !parseSearchURLQuery(props.location.search)

    const hideNavbar = isSearchHomepage

    return (
        <div className={`layout theme ${props.isLightTheme ? 'theme-light' : 'theme-dark'}`}>
            {!hideNavbar && <Navbar location={props.location} history={props.history} />}
            <Switch>
                {routes.map((route, i) => {
                    const isFullWidth = !route.forceNarrowWidth
                    const Component = route.component
                    return (
                        <Route
                            {...route}
                            key="hardcoded-key" // see https://github.com/ReactTraining/react-router/issues/4578#issuecomment-334489490
                            component={undefined}
                            // tslint:disable-next-line:jsx-no-lambda
                            render={props => (
                                <div
                                    className={[
                                        'layout__app-router-container',
                                        `layout__app-router-container--${isFullWidth ? 'full-width' : 'restricted'}`,
                                    ].join(' ')}
                                >
                                    {Component && <Component {...props} isFullWidth={isFullWidth} />}
                                    {route.render && route.render(props)}
                                </div>
                            )}
                        />
                    )
                })}
            </Switch>
        </div>
    )
}

interface AppState {
    error?: Error
    isLightTheme: boolean
}

/**
 * The root component
 */
class App extends React.Component<{}, AppState> {
    public state: AppState = {
        isLightTheme: getColorTheme() === 'light',
    }

    private subscriptions = new Subscription()

    constructor(props: {}) {
        super(props)
        // Fetch current user data
        fetchCurrentUser().subscribe(undefined, error => {
            console.error(error)
            this.setState({ error })
        })
    }

    public componentDidMount(): void {
        this.subscriptions.add(colorTheme.subscribe(theme => this.setState({ isLightTheme: theme === 'light' })))
    }

    public componentDidUpdate(): void {
        fetchCurrentUser().subscribe(undefined, error => {
            console.error(error)
            this.setState({ error })
        })
    }

    public componentWillUnmount(): void {
        this.subscriptions.unsubscribe()
    }

    public render(): JSX.Element | null {
        if (this.state.error) {
            return <HeroPage icon={ErrorIcon} title={'Something happened'} subtitle={this.state.error.message} />
        }

        if (window.pageError && window.pageError.statusCode !== 404) {
            const statusText = window.pageError.statusText
            const errorMessage = window.pageError.error
            const errorID = window.pageError.errorID

            let subtitle: JSX.Element | undefined
            if (errorID) {
                subtitle = (
                    <p>
                        Sorry, there's been a problem. Please <a href="mailto:support@sourcegraph.com">contact us</a>{' '}
                        and include the error ID:
                        <span className="error-id">{errorID}</span>
                    </p>
                )
            }
            if (errorMessage) {
                subtitle = (
                    <div className="app__error">
                        {subtitle}
                        {subtitle && <hr />}
                        <pre>{errorMessage}</pre>
                    </div>
                )
            } else {
                subtitle = <div className="app__error">{subtitle}</div>
            }
            return <HeroPage icon={ServerIcon} title={'500: ' + statusText} subtitle={subtitle} />
        }
        if (window.context.onPrem && window.context.showOnboarding) {
            return (
                <BrowserRouter>
                    <Route path="/" component={InitializePage} />
                </BrowserRouter>
            )
        }

        return (
            <BrowserRouter>
                <Route path="/" render={this.renderLayout} />
            </BrowserRouter>
        )
    }

    private renderLayout = (props: RouteComponentProps<any>) => (
        <Layout {...props} isLightTheme={this.state.isLightTheme} />
    )
}

window.addEventListener('DOMContentLoaded', () => {
    render(<App />, document.querySelector('#root'))
})

updateUserSessionStores()
