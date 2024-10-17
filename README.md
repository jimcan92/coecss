# COECSS - Class Scheduling System

The Class Scheduling System for CTU Moalboal simplifies scheduling for instructors, rooms, and sections. It supports flexible time slots and ensures efficient weekly scheduling tailored to academic requirements.

## Features

- **Instructor & Room Management**
- **Custom Time Slots**: MWF (1, 1.5 hours), TTH (2 hours)
- **Optimized for MWF & TTH Schedules**

## Technologies

- **Frontend**: [SvelteKit](https://kit.svelte.dev/), [TypeScript](https://www.typescriptlang.org/), [TailwindCSS](https://tailwindcss.com/)
- **Backend**: [Wails](https://wails.io/) (Go)
- **Database**: [bbolt](https://github.com/etcd-io/bbolt)

## Getting Started

### Prerequisites

- [Node.js](https://nodejs.org/), [Go](https://golang.org/), [Wails CLI](https://wails.io/docs/gettingstarted/installation)

### Installation

1. Clone the repo:

   ```bash
   git clone https://github.com/jimcan92/coecss.git
   cd coecss
   ```

2. Install dependencies:

   ```bash
   npm install
   ```

3. Run in development:

   ```bash
   wails dev
   ```

4. Build for production:
   ```bash
   wails build
   ```

## License

Licensed under the MIT License.
