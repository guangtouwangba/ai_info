# #vivo公开自研大模型数据筛选方法##港科大vivo数据筛选方法效率提升10倍#vivo自研大模型用的数据筛选方法，公开了。香港科技大学和vivo AI Lab联名提出PreSelect...

**URL**: https://weibo.com/6105753431/PrX7s6j3w

## 原始摘要

<a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23vivo%E5%85%AC%E5%BC%80%E8%87%AA%E7%A0%94%E5%A4%A7%E6%A8%A1%E5%9E%8B%E6%95%B0%E6%8D%AE%E7%AD%9B%E9%80%89%E6%96%B9%E6%B3%95%23&amp;extparam=%23vivo%E5%85%AC%E5%BC%80%E8%87%AA%E7%A0%94%E5%A4%A7%E6%A8%A1%E5%9E%8B%E6%95%B0%E6%8D%AE%E7%AD%9B%E9%80%89%E6%96%B9%E6%B3%95%23" data-hide=""><span class="surl-text">#vivo公开自研大模型数据筛选方法#</span></a><a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E6%B8%AF%E7%A7%91%E5%A4%A7vivo%E6%95%B0%E6%8D%AE%E7%AD%9B%E9%80%89%E6%96%B9%E6%B3%95%E6%95%88%E7%8E%87%E6%8F%90%E5%8D%8710%E5%80%8D%23&amp;extparam=%23%E6%B8%AF%E7%A7%91%E5%A4%A7vivo%E6%95%B0%E6%8D%AE%E7%AD%9B%E9%80%89%E6%96%B9%E6%B3%95%E6%95%88%E7%8E%87%E6%8F%90%E5%8D%8710%E5%80%8D%23" data-hide=""><span class="surl-text">#港科大vivo数据筛选方法效率提升10倍#</span></a><br><br>vivo自研大模型用的数据筛选方法，公开了。<br><br>香港科技大学和vivo AI Lab联名提出PreSelect，目前已被ICML 2025接收。<br><br>这是一种轻量级且高效的数据选择方法：只需要训练和部署一个基于fastText的评分器，就可以减少10倍的计算需求。<br><br>该方法提出数据的预测强度(Predictive Strength) 的概念和计算公式，利用在不同模型上Loss有序性表征数据对特定能力的贡献，通过获取特定能力的有效样本训练fastText分类器对全量训练数据进行筛选。【图1】<br><br>“压缩即智能”（compression represents intelligence）这一观点揭示了一个核心现象：大模型对数据的压缩能力（例如BPC, bits per character）与其在该数据上的归一化Loss存在等价关系，且与模型在下游任务中的表现高度相关。<br><br>换言之，模型越能高效压缩数据，模型能力或智能水平越高。<br><br>PreSelect团队提出以数据预测强度（Predictive Strength）作为衡量模型loss与下游任务（benchmark）表现一致性的指标，其计算公式如【图2】所示。<br><br>- N代表模型数量，这些模型在benchmark的得分 {S1 &lt; S2 &lt; … &lt; SN}<br>- C代表模型在数据集d上的归一化loss，即BPC<br>- Z为归一化因子<br>- I{}为指示函数<br>- S取值范围 [0,1]<br><br>当 S=1 时，表示不同模型在benchmark上的得分排序与其在该数据上的loss排序完全一致，说明该数据具有很高的预测强度；相反，当S=0时，说明两种排序之间没有相关性，该数据对下游任务的作用弱，预测强度很低。<br><br>根据预测强度的高低对数据进行筛选，优先保留那些使得不同模型在benchmark上的得分排序与在数据上的loss排序更一致的数据。<br><br>这类数据对模型能力的贡献更加显著，能够更有效地提升模型效果。<br><br>与现有方法相比，该方法具有更坚实的理论基础，减少了对人工启发规则的依赖，筛选过程更客观、更具有泛化性。<br><br>了解更多详情，欢迎点击链接：<a href="https://weibo.cn/sinaurl?u=https%3A%2F%2Fmp.weixin.qq.com%2Fs%2F3n3E0l0eKbJ-FMiEE9CC8w" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_web_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">全新预训练数据筛选方案，让数据效率提升10倍！配置仅需fastText评分器｜港科大vivo出品</span></a><img style="" src="https://tvax2.sinaimg.cn/large/006Fd7o3gy1i1gbco9fj1j30zk082762.jpg" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax1.sinaimg.cn/large/006Fd7o3gy1i1gbcq6o3ij30hk042q33.jpg" referrerpolicy="no-referrer"><br><br>

## AI 摘要

vivo与香港科技大学联合研发的PreSelect数据筛选方法被ICML 2025接收。该方法通过fastText评分器计算数据的"预测强度"(Predictive Strength)，利用模型loss与下游任务表现的排序一致性筛选高质量数据，使计算需求降低10倍。其理论基础是"压缩即智能"现象——模型压缩数据能力(BPC)与智能水平正相关。相比传统方法，PreSelect减少人工规则依赖，筛选更客观高效。实验表明，优先保留预测强度高的数据能显著提升模型效果。该成果为大模型训练提供了轻量级的数据优化方案。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-16T03:22:17Z
- **目录日期**: 2025-05-16
