# #MIT新技术让AI生成代码更准确##AI代码生成不再跑偏#代码敲烦了？一项新技术，有望让程序员们的生活变得更加轻松！来自麻省理工学院等机构的研究人员提出了一项...

**URL**: https://weibo.com/6105753431/PosQbgE05

## 原始摘要

<a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23MIT%E6%96%B0%E6%8A%80%E6%9C%AF%E8%AE%A9AI%E7%94%9F%E6%88%90%E4%BB%A3%E7%A0%81%E6%9B%B4%E5%87%86%E7%A1%AE%23&amp;extparam=%23MIT%E6%96%B0%E6%8A%80%E6%9C%AF%E8%AE%A9AI%E7%94%9F%E6%88%90%E4%BB%A3%E7%A0%81%E6%9B%B4%E5%87%86%E7%A1%AE%23" data-hide=""><span class="surl-text">#MIT新技术让AI生成代码更准确#</span></a><a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23AI%E4%BB%A3%E7%A0%81%E7%94%9F%E6%88%90%E4%B8%8D%E5%86%8D%E8%B7%91%E5%81%8F%23&amp;extparam=%23AI%E4%BB%A3%E7%A0%81%E7%94%9F%E6%88%90%E4%B8%8D%E5%86%8D%E8%B7%91%E5%81%8F%23" data-hide=""><span class="surl-text">#AI代码生成不再跑偏#</span></a><br><br>代码敲烦了？一项新技术，有望让程序员们的生活变得更加轻松！<br><br>来自麻省理工学院等机构的研究人员提出了一项新技术，能够自动指导LLM生成及符合编程语言规范，还不产生错误的文本。<br><br>在一个Python代码生成的测试任务中，采用该架构的一个小型开源模型的表现甚至是超越了一个体积是其两倍以上的专业商业闭源模型！<br><br>利用LLM来生成代码早已经不是一件新鲜事了，然而，想要真正提高效率，对生成代码的结构和语义有着严格的要求。<br><br>目前控制LLM生成结构化文本（如代码）的常见方法主要有两种，都存在相应的短板：<br><br>- 检查完整输出：确保代码有效且能无错误运行，如果不符合要求就必须重新开始，这会消耗大量计算资源；<br>- 边生成边检查：能保证代码符合编程语言规范且结构有效，但渐进式修正可能导致代码逐渐偏离用户本意，最终影响准确性。<br><br>正因此，研究人员们尝试同时对生成内容结构和语义进行双重控制。<br><br>他们提出了一种与深度学习中常见的扩展方法非常不同的方法：将专家知识编码到LLM中，引导其生成最高潜质的输出。<br><br>简单来说，他们采用了一种叫做序列蒙特卡洛的技术，让LLM并行生成多个候选结果相互竞争，模型根据各并行计算线程输出结果的质量，动态分配计算资源。<br><br>每个输出都会被赋予一个权重，代表其在结构有效性和语义准确性上的可能性。在计算的每个步骤中，模型都会聚焦于高权重的候选结果，淘汰其余结果。<br><br>从某种意义上说，这就像有一个专家在LLM背后监督，确保它在每个步骤都做出正确选择，同时始终聚焦于整体目标。<br><br>未来，该研究团队希望将该技术应用于更大规模的文本生成，而非局限于小片段处理。他们还计划将该方法与学习机制相结合，使模型在受控输出的过程中不断提升准确性。<br><br>该研究成果将在ICLR上发表。本文有三名共同第一作者，分别是麻省理工学院研究生João Loula、米拉-魁北克人工智能研究所研究员Benjamin LeBrun、约翰霍普金斯大学研究生Li Du。<br><br>论文地址：<a href="https://weibo.cn/sinaurl?u=https%3A%2F%2Fopenreview.net%2Fpdf%3Fid%3DxoXn62FzD0" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_web_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">网页链接</span></a><br>开源仓库：<a href="https://weibo.cn/sinaurl?u=https%3A%2F%2Fgithub.com%2Fgenlm%2Fgenlm-control" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_web_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">网页链接</span></a><img style="" src="https://tvax2.sinaimg.cn/large/006Fd7o3ly1i0psj7crsdj30vq0tudue.jpg" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax4.sinaimg.cn/large/006Fd7o3ly1i0psj7pubkj31fk0na12w.jpg" referrerpolicy="no-referrer"><br><br>

## AI 摘要

MIT研究人员开发了一种新技术，通过序列蒙特卡洛方法让AI并行生成多个代码候选方案，并根据结构和语义准确性动态分配计算资源。该方法将专家知识编码到大型语言模型(LLM)中，在Python代码生成测试中，小型开源模型表现优于体积更大的商业模型。相比传统方法(完整输出检查或渐进式修正)，新技术能同时保证代码规范性和语义准确性。未来计划将该技术扩展到更大规模文本生成并与学习机制结合。研究成果将在ICLR发表，相关代码已开源。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-22T22:02:46Z
- **目录日期**: 2025-04-22
