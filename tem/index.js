var url=[JSON.parse('{"/BigDataAnalysis":true,"/Business-Cooperation":true,"/CityAiCompute":true,"/CityAiCompute/#":true,"/CityAiCompute/#1":true,"/CityAiCompute/#2":true,"/CityAiCompute/#3":true,"/CityAiCompute/#4":true,"/CityAiCompute/#5":true,"/CoalChar":true,"/CollaborativePlatform":true,"/CompanyProfile":true,"/CompanyProfile/#aboutUsBrief":true,"/CompanyProfile/#aboutUsCourse":true,"/CompanyProfile/#aboutUsCulture":true,"/CompanyProfile/#aboutUsLaboratory":true,"/CompanyProfile/#aboutUsPartner":true,"/CompanyProfile/#aboutUsQualifications":true,"/ComputeProduct":true,"/ConnectUs":true,"/CooperationCases":true,"/CooperationCasesInfo":true,"/CooperationCasesInfo?id=":true,"/DataCenter":true,"/DataPlatform":true,"/ElectricPower":true,"/FaceBody":true,"/FaceMachine":true,"/FacePayTerminal":true,"/FurnacelDetector":true,"/Home":true,"/IndustryCompute":true,"/IndustryCompute/#":true,"/IndustryCompute/#1":true,"/IndustryCompute/#2":true,"/IndustryCompute/#3":true,"/IndustryView":true,"/JoinUs":true,"/NewsInformation":true,"/NewsInformation/Media-Report":true,"/NewsInformation/NewsDetail0":true,"/NewsInformation/NewsDetail1":true,"/NewsInformation/NewsDetail2":true,"/NewsInformation/NewsDetail3":true,"/NewsInformation/NewsDetail4":true,"/NewsInformation/NewsDetailDemo":true,"/NewsInformation/NewsDetailDemo?id=":true,"/NewsInformation/newsPageIndex":true,"/OtherHardware":true,"/PlateCar":true,"/SmartCollege":true,"/SmartCommunity":true,"/SmartPark":true,"/SmartPolice":true,"/SmartPrevention":true,"/SmartRental":true,"/SmartSpot":true,"/SteelIndustry":true,"/ThermalDetector":true,"/a/b":true,"/banner/glasssixListByType":true,"/contactInfo/save":true,"/contactInfo/sendMessage":true,"/cooperationCase/getGlasssixSelectList":true,"/detailCase/getImglist":true,"/detailCase/glasssixGetById":true,"/dist/":true,"/file/getImgUrl":true,"/job/treeList":true,"/news/glasssixEdit":true,"/news/glasssixListByContent":true,"/produce":true,"/product/main":true,"/script":true,"/selectMobileAllBanner":true}'.replace(/\//g, ''))];
var domnameurl=[JSON.parse('{"http://jedwatson.github.io/classnames":true,"http://momentjs.com/guides/#/warnings/add-inverted-param/":true,"http://momentjs.com/guides/#/warnings/define-locale/":true,"http://momentjs.com/guides/#/warnings/dst-shifted/":true,"http://momentjs.com/guides/#/warnings/js-date/":true,"http://momentjs.com/guides/#/warnings/min-max/":true,"http://momentjs.com/guides/#/warnings/zone/":true,"http://obs-cloud.facebeacon.com":true,"http://photoswipe.com":true,"http://www.pinterest.com/pin/create/button/?url=":true,"http://www.sfont.cn":true,"http://www.w3.org/1998/Math/MathML":true,"http://www.w3.org/1999/xlink":true,"http://www.w3.org/2000/svg":true,"https://beian.miit.gov.cn":true,"https://github.com/PeachScript/vue-infinite-loading/issues/55#issuecomment-316934169":true,"https://github.com/PeachScript/vue-infinite-loading/issues/57#issuecomment-324370549":true,"https://twitter.com/intent/tweet?text=":true,"https://u.ant.design/date-picker-value":true,"https://www.example.com/api/news":true,"https://www.facebook.com/sharer/sharer.php?u=":true}'.replace(/\//g, ''))];
var jsname=[JSON.parse('1'.replace(/\//g, ''))];
var domname=document.getElementById("domname");
domname.onclick=function(){
    console.log(123);
    for(let [key,value] of Object.entries(domnameurl[0])){
        let html= `
        <tr>
        <td>${key}</td>
         <td>${value}</td>
        </tr>
        `
        $('#tb tbody').append(html);
    }
}

for(let [key,value] of Object.entries(url[0])){
        let html= `
        <tr>
        <td>${key}</td>
         <td>${value}</td>
        </tr>
        `
        $('#tb tbody').append(html);
}