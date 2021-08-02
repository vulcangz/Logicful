import * as Sentry from "@sentry/browser";
import { Integrations } from "@sentry/tracing";

export function initSentry() {
  Sentry.init({
    dsn:
      "https://a325b288e9964eeaad6b2e2718c2d11b@o453689.ingest.sentry.io/5442669",
    integrations: [new Integrations.BrowserTracing()],
    tracesSampleRate: 1.0,
  });
}
