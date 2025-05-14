# #研究表明AI看不懂多视角画面##多模态新标准无一模型及格#多模态大语言模型的多视角理解有多难？GPT-4o竟然连及格线都没摸到。所谓“多视图”理解，是让AI像人类...

**URL**: https://weibo.com/6105753431/PrMCt7lUi

## 原始摘要

<a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E7%A0%94%E7%A9%B6%E8%A1%A8%E6%98%8EAI%E7%9C%8B%E4%B8%8D%E6%87%82%E5%A4%9A%E8%A7%86%E8%A7%92%E7%94%BB%E9%9D%A2%23&amp;extparam=%23%E7%A0%94%E7%A9%B6%E8%A1%A8%E6%98%8EAI%E7%9C%8B%E4%B8%8D%E6%87%82%E5%A4%9A%E8%A7%86%E8%A7%92%E7%94%BB%E9%9D%A2%23" data-hide=""><span class="surl-text">#研究表明AI看不懂多视角画面#</span></a><a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E5%A4%9A%E6%A8%A1%E6%80%81%E6%96%B0%E6%A0%87%E5%87%86%E6%97%A0%E4%B8%80%E6%A8%A1%E5%9E%8B%E5%8F%8A%E6%A0%BC%23&amp;extparam=%23%E5%A4%9A%E6%A8%A1%E6%80%81%E6%96%B0%E6%A0%87%E5%87%86%E6%97%A0%E4%B8%80%E6%A8%A1%E5%9E%8B%E5%8F%8A%E6%A0%BC%23" data-hide=""><span class="surl-text">#多模态新标准无一模型及格#</span></a><br><br>多模态大语言模型的多视角理解有多难？GPT-4o竟然连及格线都没摸到。<br><br>所谓“多视图”理解，是让AI像人类一样从不同角度整合信息。<br><br>此前一直缺乏评测标准，UC伯克利、香港大学、牛津等团队，联合推出All-Angles Bench基准，专门评估AI模型在遮挡、相对位置判断、相机视角识别等方面的能力。<br><br>测试覆盖90个真实场景，2100+多视图问答对，囊括Counting、Attribute Identification、Relative Distance等六大任务，全面考察AI的三维理解力。【图1】<br><br>他们评测了27个主流模型，结果发现，多模态大语言模型与人类水平之间存在显著差距。<br><br>进一步研究发现，多模态大模型存在两种主要的缺陷模式【图2】：<br><br>（1）在遮挡情况下跨视图对应能力较弱；<br>（2）对粗略相机位姿的估计能力较差。<br><br>此外，GPT-4o有时会选择每个视角中的最大数量，而不是对跨视角的物体数量进行统一统计。【图3】<br><br>通过可视化的方法，研究人员让MLLM推理多视图下的物体和相机的位置与朝向。<br><br>虽然GPT-4o和Gemini-2.0-Flash对单张图像的场景理解表现尚可，但它们在对齐不同的相机视角时存在困难，难以正确处理视角变换。【图4】<br><br>论文：<a href="https://weibo.cn/sinaurl?u=https%3A%2F%2Farxiv.org%2Fabs%2F2504.15280" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_web_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">网页链接</span></a><br>项目：<a href="https://weibo.cn/sinaurl?u=https%3A%2F%2Fdanielchyeh.github.io%2FAll-Angles-Bench%2F" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_web_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">网页链接</span></a><img style="" src="https://tvax3.sinaimg.cn/large/006Fd7o3gy1i1f0nz31pyj30z90c0k62.jpg" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax3.sinaimg.cn/large/006Fd7o3gy1i1f0omufssj30z90dtdxf.jpg" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax2.sinaimg.cn/large/006Fd7o3gy1i1f0pze66jj30zk0q21kx.jpg" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax2.sinaimg.cn/large/006Fd7o3gy1i1f0r1gugoj30zk0dh4er.jpg" referrerpolicy="no-referrer"><br><br>

## AI 摘要

UC伯克利、香港大学和牛津团队联合推出All-Angles Bench基准，评估AI模型的多视角理解能力。测试覆盖90个场景和2100多个多视角问答对，考察遮挡处理、相机视角识别等六大任务。研究发现，27个主流模型（包括GPT-4o和Gemini-2.0-Flash）表现不佳，存在两大缺陷：跨视图对应能力弱和相机位姿估计差。例如，GPT-4o会错误统计跨视角物体数量。研究表明，当前多模态大语言模型在三维理解上与人类水平仍有显著差距。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-14T07:02:51Z
- **目录日期**: 2025-05-14
