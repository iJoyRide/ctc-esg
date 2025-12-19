# High Level Diagram
<img width="1213" height="1136" alt="image" src="https://github.com/user-attachments/assets/0b4b4c66-3931-4305-b2a6-4d6fddbd26e6" />

# Setup Instructions

### Step 1: Activate the Virtual Environment
To activate the `uv` virtual environment, use the following command:

For **Windows**:
```bash
.\.venv\Scripts\activate
```

For **macOS/Linux**:
```bash
source .venv/bin/activate
```

### Step 2: Install python dependencies
To test locally and install the dependecies

```bash
uv add.
```

```bash
uv pip install -e .
```

### Step 3: Install NextJS dependencies
To test locally and install the dependecies

```bash
npm install
```

### Step 4: Build Docker Image
To build the required Docker image, including dependencies such as `uv`, run the following command:

```bash
docker compose build
```

### Step 3: Run the Docker Image
To run the image, use the following command:

```bash
docker compose up 
```
### Step 4: Go to FastAPI swagger

```bash
http://localhost:8000/docs   
```
### Alternative to running the application locally

```bash
cd backend
```
```bash
uvicorn app.main:app --host 0.0.0.0 --port 8000
```
```bash
cd frontend
```
```bash
npm run dev
```
```bash
docker compose up --build ollama
```

# Git best practices

### Step 1: Creating a new branch
```bash
git checkout -b SCRUM-7
```

### Step 2: Staging changes
```bash
git add .
```

### Step 3: Pushing changes to branch
New Features
```bash
git commit -m "feat(SCRUM-340): add .env file for environment variable management" 
```
Bug Fixes
```bash
git commit -m "fix(SCRUM-340): resolve issue with missing environment variable handling"
```
Code Refactoring
```bash
git commit -m "refactor(SCRUM-340): restructure config loader to support .env file" 
```

### Step 4: Pushing to new branch 
```bash
git push -u origin SCRUM-7
```

### Step 5: Delete local branch
```bash
git branch -d SCRUM-7
```
test








