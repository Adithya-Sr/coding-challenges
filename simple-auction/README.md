# Simple Auctioneer and Bidder Service

This project implements a simple auctioneer and bidder service, as in a challenge seen on [Anthonygg's YouTube channel](https://youtu.be/PCgnBCE-N8I?si=6uPAhIk1WdDo0DJw ). The auctioneer service receives bids from the bidder service and returns the best bid.
 

## Getting Started

### Installation

1. Clone the repository:
    ```bash
    git clone <repository_url>
    ```

2. Navigate to the `simple-auction` directory:
    ```bash
    cd simple-auction
    ```

### Usage

To run the project, follow these steps:

1. Run the services using `make run`. This command will start both the auctioneer and bidder services:
    ```bash
    make run
    ```

2. After running the services, you can make requests to the auctioneer service. Use the following endpoint:
    ```
    http://localhost:3000/AdPlacement?placementId={any_integer}
    
    ```
    Replace `{any_integer}` with an integer value representing the placement ID.

### Testing

To run tests for the project, use the following command:
```bash
make test
