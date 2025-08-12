# Resume Maker Frontend

A modern Vue.js frontend for the Resume Maker application that connects to the Go backend API.

## Features

- **Personal Information Management**: Edit name, description, email, phone, and address
- **Work Experience**: Add, edit, and delete work experience entries
- **Skills Management**: Add skills with proficiency levels (0-100%)
- **API Key Authentication**: Secure editing with API key authentication
- **Responsive Design**: Modern UI built with Tailwind CSS
- **Real-time Updates**: Instant updates when data changes

## Tech Stack

- **Vue 3** with Composition API
- **TypeScript** for type safety
- **Pinia** for state management
- **Vue Router** for navigation
- **Tailwind CSS** for styling
- **Vite** for build tooling

## Prerequisites

- Node.js 18+ 
- npm or yarn

## Setup

1. **Install dependencies**:
   ```bash
   npm install
   ```

2. **Start the development server**:
   ```bash
   npm run dev
   ```

3. **Open your browser** and navigate to `http://localhost:5173`

## Usage

### Viewing Resume Data
- The application automatically loads your resume data from the backend
- Personal information, experience, and skills are displayed in a clean, professional layout

### Editing Resume Data
1. **Set API Key**: Enter your API key in the header to enable editing
2. **Edit Personal Info**: Click the "Edit" button in the Personal Information section
3. **Add Experience**: Click "Add Experience" to add new work experience
4. **Add Skills**: Click "Add Skill" to add new skills with proficiency levels
5. **Edit/Delete**: Use the edit and delete buttons on each item (when API key is set)


## Development

### Available Scripts

- `npm run dev` - Start development server
- `npm run build` - Build for production
- `npm run preview` - Preview production build
- `npm run lint` - Run ESLint
- `npm run format` - Format code with Prettier

### API Endpoints

The frontend connects to these backend endpoints:
- `GET /api/personal-info` - Fetch personal information
- `PUT /api/personal-info` - Update personal information
- `GET /api/experience/` - Fetch work experience
- `POST /api/experience/` - Add new experience
- `PUT /api/experience/` - Update experience
- `DELETE /api/experience/{id}` - Delete experience
- `GET /api/skills/` - Fetch skills
- `POST /api/skills/` - Add new skill
- `PUT /api/skills/` - Update skill
- `DELETE /api/skills/{id}` - Delete skill

## Building for Production

```bash
npm run build
```

The built files will be in the `dist/` directory, ready for deployment.

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Run tests and linting
5. Submit a pull request

## License

This project is licensed under the Apache License 2.0.
