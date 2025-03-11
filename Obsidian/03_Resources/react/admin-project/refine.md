## Areas

# Using Refine to build Admin Dashboards.

## Providers

These are the building blocks of a Refine application, it provides things like

- Data Fetching
- Auth
- Routing
- Notifications
- Real Time Updates

### Data Providers

> [!IMPORTANT]
> Check out API Consuming Flow

We use providers to get data to render our stuff, it connects to a datastore which we can use to get data.

- Create providers that will get data from our GraphQL backend.
- Fetch Wrapper: Helps us with making custom calls to API, providing headers and even handling errors.
- GetGraphQLClinet returns Apollo graphQL I guess.
- wsClient does something too.
- data provider
- Live provider
- Checkout Live Provider:
  - GraphQL-ws: This will listen to changes made to this graphQL API.
- Auth Providers: We will build our own.
- TypeScript workflow setup. (so we don't have to write types)
- Create Routes

> [!NOTE]
> Where is the data coming from ?
