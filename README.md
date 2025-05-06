# 📊 Hypothesis Test Data Analysis (Go)

This project tests a real-life hypothesis using basic statistical analysis in Go:

> **Hypothesis:** *"Fiction books are cheaper than Non Fiction books."*

We perform a z-test on book price data from Kaggle to confirm or reject this hypothesis using real-world data.

---

## 🗂 Project Structure
````
HypothesisTestDataAnalysis/
├── cmd/
│   └── main.go       
├── internal/
│   └── analysis.go 
├── bestsellers with categories.csv
├── go.mod
└── README.md

````

---

## 📦 Prerequisites

- Go 1.18 or higher
- Git (optional)
- Dataset: [Amazon Top 50 Bestselling Books 2009–2019](https://www.kaggle.com/datasets/sootersaalu/amazon-top-50-bestselling-books-2009-2019)

---

## 🔧 Setup

1. **Clone the project** (or download ZIP):

```bash
git clone https://github.com/meruyert4/HypothesisTestDataAnalysis.git
cd HypothesisTestDataAnalysis
````

## 🚀 Run the Project

```bash
go run cmd/main.go
```

You should see output like:

```
📚 Fiction Avg Price: $10.85 (240 books)
📘 Non Fiction Avg Price: $14.84 (310 books)
Z-score: -4.55 → Reject H₀. Fiction books are significantly cheaper. ✅
```

---

## 📚 Dataset Citation

Sootla, S. (2020). *Amazon Top 50 Bestselling Books 2009–2019* \[Data set]. Kaggle.
🔗 [https://www.kaggle.com/datasets/sootersaalu/amazon-top-50-bestselling-books-2009-2019](https://www.kaggle.com/datasets/sootersaalu/amazon-top-50-bestselling-books-2009-2019)

---

## 📌 Notes

* This project uses **population standard deviation**.
* Z-score interpretation uses a **one-tailed test** at 95% confidence (threshold: -1.645).

---