Architecture choices
--------------------
This tool is structured as small, focused packages: `types` for data shapes, `utils` for helpers, and a `main` entry that coordinates I/O and AI calls. This pattern which obeys single-responsibility principle  keeps serialization, URL-generation, and AI wiring decoupled so each piece can be tested or replaced independently.

Concurrency strategy
--------------------
Processing is synchronous and sequential with a rate limiter (fixed interval). This simplifies ordering, avoids state races, and respects API rate limits. This also favors predictability and easier local debugging.

Error handling approach
-----------------------
The dominant error handling approach is log and continue. When a critical error occurs during the processing/evaluation of a single struct item, the error is logged, and the loop continues processing the remaining items, thereby preventing a single bad item from stopping the entire batch job.
    The program also uses fail-fast for critical startup errors (config, client init, file I/O) as this errors prevents the core logic from running correctly. 
    Network/API calls include sensible timeouts and can be retried externally; adding an exponential backoff retry wrapper around transient network errors is straightforward if higher reliability is required.

Performance vs safety trade-offs
------------------------------
Sequential processing of tweets is slow because it will wait 6 seconds even if the API responds in 1 sec.

Language & framework choice
-------------------------
I am currently exploring Go so i chose Go to improve my Go writing skills.