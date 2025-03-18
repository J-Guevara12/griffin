# Griffin
**Gri**nding **Fin**nish is a personal project to develop a task management tool to keep myself organized, will be written in GoLang and I expect to use it in desktop environments and mobile devices.

## Working plan:

### Phase 1:
The data will be stored in MongoDB (initially self hosted) and the idea is to interact with it using a CLI (Cobra):

- [X] Create the database and set up the initial schema (required fields only)
- [X] Set up functions to perform CRUD operations on the database
- [ ] Bind these functions with CLI commands
- [ ] Add options to the `ls` command (sort, filter, output format, limit).
- [ ] Add `created`, `closed` and `modified` fields.
- [ ] Add configuration options (.conf, .yaml, .env)
- [ ] Dependency injection over the `db_writer()` object.
- [X] Visualize short fields (Summary, due date, priority) in a tabular format.
- [ ] Add Markdown support for the descriptions.
- [ ] Enable custom fields.
- [ ] Handle Status and Priority values as database objects (inside a collection).
- [ ] Perform CRUD operations over the previous fields (including custom).
- [ ] Enable value-limited (like a select) fields.
- [ ] Color palette via configuration file (Include previsualizer).

### Phase 2:
This phase involves the development of a TUI able to interact with the system and facilitate the CLI operations and adding a few cool features:
- [ ] Select any task and open it.
- [ ] Task creation form.
- [ ] Task update form (based on the previous).
- [ ] Task deletion from the TUI.
- [ ] Task Bulk operations:
    - [ ] Delete tasks
    - [ ] Update the same field in multiple tasks
- [ ] Group tasks by workspaces.
- [ ] Create project objects that contain multiple tasks and a broader set of fields (Notes, links, actors).
- [ ] Promote tasks to projects.
- [ ] Improved user experience to select dates.
- [ ] Custom field editing UI.
- [ ] Fuzzy finder (tentative).

### Phase 3:
I want to be able to remember any of the tasks I have to work and other activities, the idea here is to build a daemon capable of notifying me.
- [ ] Set up a calendar object (containing tasks and events)
- [ ] Google Calendar integration
- [ ] OS Notification.
- [ ] Messaging notification:
    - [ ] Whatsapp
    - [ ] Email
    - [ ] Slack
    - [ ] Teams
    - [ ] Telegram
- [ ] Configure which type of tasks are related to certain notificator.
- [ ] Custom notifications for each task (message, priority, times it will notify).
- [ ] Set up the daemon (should write an aknowlegedment to the database).

### Phase 4:
I will probably want to access my system from multiple devices, a remote database and backend will have to be developed:
- [ ] Implement a backend that performs the *Phase 1* Operations (CRUD on the database) as an HTTP server.
- [ ] Implement authentication and authorization (SPIFFE library seems an interesting option).
- [ ] Modify the CLI and TUI to perform API calls instead of direct database operations.
- [ ] When interacting from a mobile device, instead of writing a full mobile-app I'd preffer to develop a telegram/whatsapp bot that allows me to query my most recent tasks and perform basic update operations (silence notifications for instance).

In this phase some kind of observability, monitoring will have to be done, my first option would be to use Prometheus + Graphana, additionally I plan to use a GitOps approach to the deployment of the backend using docker containers and Kubernetes (maybe).

## Repositories:
The code will be organized in a CLI + TUI repo and a sepparate one for the web backend.

### CLI + TUI:
Initially will be the only one and will work with a local MongoDB deployment on docker.

Will have a Github Actions CI/CD Pipeline that builds, tests and publishes a release of the app for the following platforms:
* Windows (x86)
* Macos (ARM)
* Linux (AMD64)

When the *Phase 3* is finished, I will leave the purely local implementation as a sepparate branch for anyone who wants to use it (or I may have a flag on the CLI that indicates if the database is local).

## Backend
Will run have the backend which will be a docker container in Golang.

Will have a Github Actions CI/CD pipeline that builds, test and publishes the image that will be polled by the production server. I intend to run it in a Linux VM in a yet to define VPS.
