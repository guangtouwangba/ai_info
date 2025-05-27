# #字节复旦团队推出自适应推理框架##摆脱模型对思维链过度依赖#过度依赖CoT思维链推理会降低模型性能，有新解了！来自字节、复旦大学的研究人员提出自适应推理框...

**URL**: https://weibo.com/6105753431/PtLjhcgtn

## 原始摘要

<a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E5%AD%97%E8%8A%82%E5%A4%8D%E6%97%A6%E5%9B%A2%E9%98%9F%E6%8E%A8%E5%87%BA%E8%87%AA%E9%80%82%E5%BA%94%E6%8E%A8%E7%90%86%E6%A1%86%E6%9E%B6%23&amp;extparam=%23%E5%AD%97%E8%8A%82%E5%A4%8D%E6%97%A6%E5%9B%A2%E9%98%9F%E6%8E%A8%E5%87%BA%E8%87%AA%E9%80%82%E5%BA%94%E6%8E%A8%E7%90%86%E6%A1%86%E6%9E%B6%23" data-hide=""><span class="surl-text">#字节复旦团队推出自适应推理框架#</span></a><a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E6%91%86%E8%84%B1%E6%A8%A1%E5%9E%8B%E5%AF%B9%E6%80%9D%E7%BB%B4%E9%93%BE%E8%BF%87%E5%BA%A6%E4%BE%9D%E8%B5%96%23&amp;extparam=%23%E6%91%86%E8%84%B1%E6%A8%A1%E5%9E%8B%E5%AF%B9%E6%80%9D%E7%BB%B4%E9%93%BE%E8%BF%87%E5%BA%A6%E4%BE%9D%E8%B5%96%23" data-hide=""><span class="surl-text">#摆脱模型对思维链过度依赖#</span></a><br><br>过度依赖CoT思维链推理会降低模型性能，有新解了！<br><br>来自字节、复旦大学的研究人员提出自适应推理框架CAR，能根据模型困惑度动态选择短回答或详细的长文本推理，最终实现了准确性与效率的最佳平衡。<br><br>推理能力的进步极大提升了大语言模型（LLMs）和多模态大语言模型（MLLMs）在各类任务中的表现。<br><br>但已有研究发现，长CoT推理并非总能提升准确率，甚至会削弱模型处理简单任务的能力（可能产生冗长输出）。<br><br>为此，研究人员提出了CAR这一基于置信度的自适应推理框架，它首先生成简短回答并评估困惑度，仅在模型置信度低（困惑度高）时触发推理。<br><br>在多模态视觉问答、关键信息提取及文本推理等多个基准测试中，CAR超越了单纯的短回答与长推理方法，在准确性与效率之间取得了最佳平衡。 <a href="https://weibo.com/ttarticle/p/show?id=2309405170933752529137" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_article_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">低Token高精度！字节复旦推出自适应推理框架CAR</span></a><img style="" src="https://tvax1.sinaimg.cn/large/006Fd7o3gy1i1u2dr7qz3j30mg0cnq4p.jpg" referrerpolicy="no-referrer"><br><br>

## AI 摘要

字节跳动与复旦大学团队提出自适应推理框架CAR，通过动态选择短回答或长推理来优化模型性能。研究发现传统思维链(CoT)推理可能降低简单任务准确率。CAR首先生成简短回答并评估困惑度，仅在置信度低时触发详细推理。在视觉问答、信息提取等测试中，CAR在准确性和效率上均优于纯短回答或长推理方法，实现了最佳平衡。该框架减少了模型对冗长推理的依赖，提升了处理效率。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-27T09:04:23Z
- **目录日期**: 2025-05-27
