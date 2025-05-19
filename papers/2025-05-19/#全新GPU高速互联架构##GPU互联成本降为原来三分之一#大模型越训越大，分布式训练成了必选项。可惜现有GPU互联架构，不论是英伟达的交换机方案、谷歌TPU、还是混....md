# #全新GPU高速互联架构##GPU互联成本降为原来三分之一#大模型越训越大，分布式训练成了必选项。可惜现有GPU互联架构，不论是英伟达的交换机方案、谷歌TPU、还是混...

**URL**: https://weibo.com/6105753431/Psy4ahH5W

## 原始摘要

<a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E5%85%A8%E6%96%B0GPU%E9%AB%98%E9%80%9F%E4%BA%92%E8%81%94%E6%9E%B6%E6%9E%84%23&amp;extparam=%23%E5%85%A8%E6%96%B0GPU%E9%AB%98%E9%80%9F%E4%BA%92%E8%81%94%E6%9E%B6%E6%9E%84%23" data-hide=""><span class="surl-text">#全新GPU高速互联架构#</span></a><a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23GPU%E4%BA%92%E8%81%94%E6%88%90%E6%9C%AC%E9%99%8D%E4%B8%BA%E5%8E%9F%E6%9D%A5%E4%B8%89%E5%88%86%E4%B9%8B%E4%B8%80%23&amp;extparam=%23GPU%E4%BA%92%E8%81%94%E6%88%90%E6%9C%AC%E9%99%8D%E4%B8%BA%E5%8E%9F%E6%9D%A5%E4%B8%89%E5%88%86%E4%B9%8B%E4%B8%80%23" data-hide=""><span class="surl-text">#GPU互联成本降为原来三分之一#</span></a><br><br>大模型越训越大，分布式训练成了必选项。可惜现有GPU互联架构，不论是英伟达的交换机方案、谷歌TPU、还是混合型TPUv4，都面临高成本、难扩展、易故障等问题。<br><br>对此，北京大学、阶跃星辰和曦智科技推出InfiniteHBD，这是一种以低成本光交换模组为核心的新一代高带宽域架构。<br><br>该架构的核心在于：<br><br>- 利用硅光子技术，把光交换能力塞进GPU间的光电模组，实现灵活、低延迟的互联。<br><br>- 动态构建“可重配置环状拓扑”，任一节点故障也能快速绕过，训练不中断。<br><br>- 加上自研的HBD-DCN编排算法，TP分布更合理，DCN压力也更小。<br><br>成本方面，InfiniteHBD比NVL-72便宜69%，能耗也更低。仿真数据和实测结果都显示，GPU浪费率几乎为零，MFU比传统DGX系统高3.37倍。<br><br>项目论文已被SIGCOMM 2025接收，背后是一支兼具学术与产业力量的团队。<br><br>该技术能否成为大模型训练的下一代标配？我们一起来看：<a href="https://weibo.cn/sinaurl?u=https%3A%2F%2Fmp.weixin.qq.com%2Fs%2FNF9hVDEanbvIEuansGDbiQ" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_web_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">全新GPU高速互联设计，为大模型训练降本增效！北大/阶跃/曦智提出新一代高带宽域架构</span></a><img style="" src="https://tvax2.sinaimg.cn/large/006Fd7o3gy1i1kugkgqomj30zk0cgaem.jpg" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax1.sinaimg.cn/large/006Fd7o3gy1i1kugklqpsj30ze0k0qb7.jpg" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax4.sinaimg.cn/large/006Fd7o3gy1i1kugkhqakj30zk0heaip.jpg" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax2.sinaimg.cn/large/006Fd7o3gy1i1kugkkgn4j30zk0jb7g0.jpg" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax4.sinaimg.cn/large/006Fd7o3gy1i1kugjvbr7j30zk0acqd6.jpg" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax1.sinaimg.cn/large/006Fd7o3gy1i1kugktg40j30zk0hdtng.jpg" referrerpolicy="no-referrer"><br><br>

## AI 摘要

北京大学、阶跃星辰和曦智科技联合推出新型GPU高速互联架构InfiniteHBD，通过硅光子技术将光交换模组集成到GPU间，实现低成本、高灵活性的互联。该架构采用可重配置环状拓扑，能快速绕过故障节点，结合自研HBD-DCN算法优化负载分布。相比英伟达NVL-72方案，成本降低69%，能耗更优，GPU利用率接近100%，模型训练效率提升3.37倍。相关论文已被SIGCOMM 2025接收，有望成为大模型训练的下一代高效解决方案。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-19T11:04:29Z
- **目录日期**: 2025-05-19
