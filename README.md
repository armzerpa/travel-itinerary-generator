# Travel Itinerary Generator API

This project is a Golang API that generates travel itineraries using the Gemini Large Language Model (LLM). It demonstrates the power of combining AI-assisted development with external AI services to create useful applications.


## Features

- Generate travel itineraries based on user input
- Integration with Gemini LLM for content generation
- RESTful API design
- Input validation
- Unit tests

## Prerequisites

- Go 1.16 or higher
- Gemini LLM API key

## Installation

1. Clone the repository:

git clone https://github.com/your-username/travel-itinerary-generator.git

2. Navigate to the project directory:

cd travel-itinerary-generator

3. Install dependencies:

go mod tidy

4. Set up your Gemini LLM API key as an environment variable:

export API_KEY=your_api_key_here

## Usage

To start the server, run:

go run main.go

The server will start on `http://localhost:8080`.

## API Endpoints

### Generate Itinerary

- **URL**: `/generate`
- **Method**: `POST`
- **Body**:
  ```json
  {
    "prompt": "4-day trip to Rome",
    "city": "Rome",
    "numberOfLines": 10,
    "type": "text"
  }

RESPONSE:
{
  "message": "Generated itinerary here..."
}


This README provides a comprehensive overview of your project, including how to set it up, use it, and understand its structure. It also highlights the unique aspect of using AI in the development process, which could be interesting to potential contributors or users of your API.

Feel free to adjust any parts of this README to better fit your project's specifics or your personal preferences.

