package model

//MovieInfo ...
type MovieInfo struct {
	MovieID      int     `gorm:"primary_key;AUTO_INCREMENT;comment:'电影ID';not null" json:"MovieID" form:"MovieID"`
	MovieName    string  `gorm:"type:varchar(20);comment:'电影名称';not null" json:"MovieName" form:"MovieName"`
	Director     string  `gorm:"type:varchar(20);comment:'导演';not null" json:"Director" form:"Director"`
	Starring     string  `gorm:"comment:'主演';not null" json:"Starring" form:"Starring"`
	Region       string  `gorm:"type:varchar(10);comment:'地区';not null" json:"Region" form:"Region"`
	MovieTypes   string  `gorm:"type:varchar(10);comment:'电影类型';not null" json:"MovieTypes" form:"MovieTypes"`
	WantNumber   int     `gorm:"default:0;comment:'想看人数';not null" json:"WantNumber"`
	MovieLength  int     `gorm:"type:smallint(2);comment:'电影时长';not null" json:"MovieLength" form:"MovieLength"`
	ReleaseTime  string  `gorm:"type:varchar(20);comment:'上映时间'" json:"ReleaseTime" form:"ReleaseTime"`
	Status       int     `gorm:"type:tinyint(1);comment:'电影状态';not null" json:"Status" form:"Status"`
	Score        float32 `gorm:"type:float(3,1);default:0.0;comment:'电影评分'" json:"Score"`
	Total        int     `gorm:"default:0;comment:'评分人数'" json:"Total"`
	BoxOffice    int     `gorm:"type:bigint;default:0;comment:'票房';not null" json:"BoxOffice"`
	Introduction string  `gorm:"type:varchar(1024);comment:'电影简介';not null" json:"Introduction" form:"Introduction"`
	Poster       string  `gorm:"comment:'电影海报';not null" json:"Poster"`
	Time         `json:"-"`
	//ScoreInfo *ScoreInfo `gorm:"ForeignKey:MovieID" json:"-"`
}
