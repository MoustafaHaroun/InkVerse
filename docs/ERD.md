```mermaid
---
title: Inkverse ERD ver 0.0.1
---

erDiagram


Novel {
    Id UUID "PK"
    AuthorId UUID "FK"
    Title string "U"
    Synopsis string  
    Rating Double
    CreatedAt string  
    ModifiedAt string
}

Chapter {
    Id UUID "PK"
    NovelId UUID "FK"
    Title string
    Content string
    CreatedAt string  
    ModifiedAt string
}

Tag {
    Id UUID "PK"
    Name string "U" 
}

Genre {
    Id UUID "PK"
    Name string "U" 
}

Review {
    Id UUID "PK"
    UserId UUID "FK"
    NovelId UUID "FK"
    Rating Double
    Status string "NN"
    Content string 
    CreatedAt string  
    ModifiedAt string
}

Comment {
    Id UUID "PK"
    UserId UUID "FK"
    ChapterId UUID "FK"
    Content string 
}


Novel ||--o{ Chapter : "n-has-multiple-c"
Novel ||--o{ Review : "n-has-multiple-r"
Novel }o--o{ Tag : "n-has-multiple-t"
Novel }o--o{ Genre : "n-has-multiple-g"

Chapter ||--o{ Comment : "c-has-multiple-c"

```