package model

type RakutenBookAPIResponse struct {
    GenreInformation []interface{} `json:"GenreInformation"`
    Items            []BookItem    `json:"Items"`
    Carrier          int           `json:"carrier"`
    Count            int           `json:"count"`
    First            int           `json:"first"`
    Hits             int           `json:"hits"`
    Last             int           `json:"last"`
    Page             int           `json:"page"`
    PageCount        int           `json:"pageCount"`
}

type BookItem struct {
    Item BookData `json:"Item"`
}

type BookData struct {
    AffiliateUrl   string `json:"affiliateUrl"`
    ArtistName     string `json:"artistName"`
    Author         string `json:"author"`
    Availability   string `json:"availability"`
    BooksGenreId   string `json:"booksGenreId"`
    ChirayomiUrl   string `json:"chirayomiUrl"`
    DiscountPrice  int    `json:"discountPrice"`
    DiscountRate   int    `json:"discountRate"`
    Hardware       string `json:"hardware"`
    Isbn           string `json:"isbn"`
    ItemCaption    string `json:"itemCaption"`
    ItemPrice      int    `json:"itemPrice"`
    ItemUrl        string `json:"itemUrl"`
    Jan            string `json:"jan"`
    Label          string `json:"label"`
    LargeImageUrl  string `json:"largeImageUrl"`
    LimitedFlag    int    `json:"limitedFlag"`
    ListPrice      int    `json:"listPrice"`
    MediumImageUrl string `json:"mediumImageUrl"`
    Os             string `json:"os"`
    PostageFlag    int    `json:"postageFlag"`
    PublisherName  string `json:"publisherName"`
    ReviewAverage  string `json:"reviewAverage"`
    ReviewCount    int    `json:"reviewCount"`
    SalesDate      string `json:"salesDate"`
    SmallImageUrl  string `json:"smallImageUrl"`
    Title          string `json:"title"`
}
