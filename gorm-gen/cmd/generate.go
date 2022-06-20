package main

import (
	"github.com/golang/mock/mockgen/model"
	"gorm.io/gen"
)

func main() {
	// 指定生成代码的具体(相对)目录，默认为：./query
	// 默认情况下需要使用WithContext之后才可以查询，但可以通过设置gen.WithoutContext避免这个操作
	g := gen.NewGenerator(gen.Config{
		// 最终package不能设置为model，在有数据库表同步的情况下会产生冲突，若一定要使用可以单独指定model package的新名字
		OutPath:      "../dal/query",
		ModelPkgPath: "../dal/model", // 默认情况下会跟随OutPath参数，在同目录下生成model目录
		/* Mode: gen.WithoutContext,*/
	})

	// 复用工程原本使用的SQL连接配置db(*gorm.DB)
	// 非必需，但如果需要复用连接时的gorm.Config或需要连接数据库同步表信息则必须设置
	g.UseDB(db)

	peopleTbl := g.GenerateModelAs("people", "People") // 指定对应表格的结构体名称
	// 为指定的结构体或表格生成基础CRUD查询方法，ApplyInterface生成效果的子集
	g.ApplyBasic(
		model.User{},
		peopleTbl,
	)

	// 为指定的数据库表实现除基础方法外的相关方法, 同时也会生成ApplyBasic对应的基础方法
	// 可以认为ApplyInterface方法是ApplyBasic的扩展版
	g.ApplyInterface(func(model.SearchByTenantMethod, model.UpdateByTenantMethod) {}, // 指定方法interface，可指定多个
		model.Order{},
		g.GenerateModel("Company"), // 在这里调用也会生成ApplyBasic对应的基础方法
	)

	// 执行并生成代码
	g.Execute()
}
