# source structure

- `src`
  - `monitor`: the WEB UI
  - `monitor-backend`: the WEB UI backend (implements the event pusher and the endpoints for the UI)
  - `rule-evaluator`: the component evaluating the user defined rules and triggering critics
  - `data-collector`: collects data from the API-Server on demand on behalf of the rule-evaluator
