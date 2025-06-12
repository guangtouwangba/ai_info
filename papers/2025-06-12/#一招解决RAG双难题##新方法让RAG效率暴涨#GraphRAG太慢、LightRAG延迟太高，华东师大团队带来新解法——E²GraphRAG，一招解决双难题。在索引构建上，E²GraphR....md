# #一招解决RAG双难题##新方法让RAG效率暴涨#GraphRAG太慢、LightRAG延迟太高，华东师大团队带来新解法——E²GraphRAG，一招解决双难题。在索引构建上，E²GraphR...

**URL**: https://weibo.com/6105753431/PwcKb1LBv

## 原始摘要

<a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E4%B8%80%E6%8B%9B%E8%A7%A3%E5%86%B3RAG%E5%8F%8C%E9%9A%BE%E9%A2%98%23&amp;extparam=%23%E4%B8%80%E6%8B%9B%E8%A7%A3%E5%86%B3RAG%E5%8F%8C%E9%9A%BE%E9%A2%98%23" data-hide=""><span class="surl-text">#一招解决RAG双难题#</span></a><a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E6%96%B0%E6%96%B9%E6%B3%95%E8%AE%A9RAG%E6%95%88%E7%8E%87%E6%9A%B4%E6%B6%A8%23&amp;extparam=%23%E6%96%B0%E6%96%B9%E6%B3%95%E8%AE%A9RAG%E6%95%88%E7%8E%87%E6%9A%B4%E6%B6%A8%23" data-hide=""><span class="surl-text">#新方法让RAG效率暴涨#</span></a><br><br>GraphRAG太慢、LightRAG延迟太高，华东师大团队带来新解法——E²GraphRAG，一招解决双难题。<br><br>在索引构建上，E²GraphRAG比GraphRAG快10倍，查询速度更是LightRAG的1/100，效率大幅提升，同时保留信息处理的细粒度和准确度。<br><br>传统RAG侧重向量检索，难以处理全局信息。GraphRAG借助大模型构图提升粒度理解，但调用成本高。LightRAG试图减少大模型调用，但仍有不少开销。<br><br>而E²GraphRAG采用两步走：<br><br>1. 用大模型总结文档树，逐层提取不同粒度的信息；<br><br>2. 用SpaCy抽实体，基于句内共现构图，提取实体间关系。<br><br>这样既有总结树提供信息层级，又有实体图增强检索精准度，构建出一个轻量但信息丰富的检索系统。 <a href="https://weibo.com/ttarticle/p/show?id=2309405176754871861385" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_article_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">GraphRAG太慢LightRAG延迟高？华东师大新方法一招破解双重难题</span></a><img style="" src="https://tvax4.sinaimg.cn/large/006Fd7o3gy1i2cmxzfsb9j30rn0fkdij.jpg" referrerpolicy="no-referrer"><br><br>

## AI 摘要

华东师范大学团队提出E²GraphRAG新方法，有效解决GraphRAG速度慢和LightRAG延迟高的问题。该方法在索引构建上比GraphRAG快10倍，查询速度是LightRAG的100倍。E²GraphRAG采用两步策略：先用大模型构建文档总结树，提取多粒度信息；再用SpaCy抽取实体关系图。这种设计既保留了信息的细粒度，又显著提升了检索效率，实现了轻量级但高精度的检索系统。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-12T22:03:20Z
- **目录日期**: 2025-06-12
